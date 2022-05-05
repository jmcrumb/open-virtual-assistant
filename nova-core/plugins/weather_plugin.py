from core.abstract_plugin import NovaPlugin
from requests import get
from json import loads
import string

class WeatherPlugin(NovaPlugin):

    def __init__(self):
        self.home = 'San Diego'
        self.last_location = None
        self.key = ''

    def get_keywords(self) -> list:
        return ['weather', 'temperature']

    def _parse_current_json(self, json: dict) -> str:
        return ('The weather in {} is currently {} degrees and {}.'.format(json['location']['name'],
            round(json['current']['temp_f']), json['current']['condition']['text']))

    def _parse_current_json_full(self, json: dict) -> str:
        return ('The weather in {} is currently {} degrees and {}. There are {} mile per hour winds'
            ' with ocational gusts of {} miles per hour. The humidity today is at {} percent, and'
            ' the air quality has scored a {} on the US EPA Index'.format(json['location']['name'],
            round(json['current']['temp_f']), json['current']['condition']['text'], 
            round(json['current']['wind_mph']), round(json['current']['gust_mph']), round(json['current']['humidity']), 
            json['current']['air_quality']['us-epa-index']))

    def _parse_forecast_json(self, json: dict) -> str:
        msg = ('Here is the upcoming forcast for {}. Today, you can expect a high of {} and a low'
            ' of {} degrees with wind speeds around {}. Tomorrow, it will be as warm as {} and as'
            ' cool as {} with {} mile per hour winds.'.format(json['location']['name'], 
            round(json['forecast']['forecastday'][0]['day']['maxtemp_f']), round(json['forecast']['forecastday'][0]['day']['mintemp_f']), 
            round(json['forecast']['forecastday'][0]['day']['maxwind_mph']), round(json['forecast']['forecastday'][1]['day']['maxtemp_f']), 
            round(json['forecast']['forecastday'][1]['day']['mintemp_f']), round(json['forecast']['forecastday'][1]['day']['maxwind_mph'])))

        if (len(json['forecast']['forecastday']) > 1):
            avg_high = 0
            avg_low = 0
            avg_wind = 0
            rainy_days = []
            for day in json['forecast']['forecastday']:
                avg_high += day['day']['maxtemp_f']
                avg_low += day['day']['mintemp_f']
                avg_wind += day['day']['maxwind_mph']
                if day['day']['daily_will_it_rain'] > 0:
                    rainy_days.append(day['date'])

            avg_high = round(avg_high / len(json['forecast']['forecastday']))
            avg_low = round(avg_low / len(json['forecast']['forecastday']))
            avg_wind = round(avg_wind / len(json['forecast']['forecastday']))

            msg += ('Later this week, you\'ll see highs of {} degrees and lows of {} degrees on'
            ' average. Winds will be about {} miles per hour.'.format(avg_high, avg_low, avg_wind))

        return msg

    def _get_current(self, location: str, full: bool) -> str:
        api_call = ('http://api.weatherapi.com/v1/current.json?key={}'
            '&q={}'.format(self.key, location))
        if full:
            rsp = get(api_call + '&aqi=yes')

            if rsp.ok:
                msg = self._parse_current_json_full(loads(rsp.text))
            else:
                jsRsp = rsp.json()
                msg = ('There seems to be an issue with the weather API. A {} error as recieved.\n'
                    'Error Code: {}'.format(rsp.status_code, jsRsp['error']['code']))
        else:
            rsp = get(api_call + '&aqi=no')
            msg = self._parse_current_json(loads(rsp.text))

        return msg

    def _get_forecast(self, location: str, days: int) -> str:
        api_call = ('http://api.weatherapi.com/v1/forecast.json?key={}'
            '&q={}&days={}&aqi=no&alerts=no'.format(self.key, location, days))
        rsp = get(api_call)

        if rsp.ok:
                msg = self._parse_forecast_json(loads(rsp.text))
        else:
            jsRsp = rsp.json()
            msg = ('There seems to be an issue with the weather API. A {} error as recieved.\n'
                'Error Code: {}'.format(rsp.status_code, jsRsp['error']['code']))

        return msg

    def _get_weather(self, command: str, location: str) -> str:
        if 'in' in command:
            location = command[command.rfind("in")+2:].strip()
        if 'tomorrow' in command:
            msg = self._get_forecast(location, 1)
        elif 'week' in command:
            msg = self._get_forecast(location, 7)
        elif 'forcast' in command or 'few days' in command:
            msg = self._get_forecast(location, 3)
        elif 'weather' in command or 'temperature' in command:
            msg = self._get_current(location, ('full' in command))
        else:
            msg = None
        return msg

    def _set_home(self, command) ->str:
        if 'to' in command:
            self.home = command[command.rfind("to")+2:].strip()
            msg = 'Your home location has been set to {}.'.format(self.home)
        else:
            msg = ('When setting a new home location please say,'
                'Nova, set my home in the weather plugin to San Diego.')
        return msg

    def execute(self, command: str) -> str:
        location = self.home
        print('\n{}'.format(command))
        if ('how' in command or 'what' in command or 'can'in command):
            msg = self._get_weather(command, location)
        elif ('set' in command):
            msg = self._set_home(command)
        else:
            msg = None
        return msg

    def help_command(self, command: str) -> str:
        return ('This weather plugin is powered by WeatherAPI.com.'
            ' It is an application that can help you quickly get the local weather and other conditons.'
            ' To get the weather, ask:'
            ' How\'s the weather or what\'s the weather like today?'
            ' To get the weather of a specific location, ask:'
            ' How\'s the weather in San Diego or what\'s the weather like in Seattle?')