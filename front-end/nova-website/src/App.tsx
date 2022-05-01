import Sandbox from "./components/sandbox";
import React from "react";
import "./styles.scss";
import "./components/PluginPreview.css"
import "./components/Rating.scss"
import { ReactQueryDevtools } from "react-query/devtools";
import { QueryClient, QueryClientProvider} from "react-query";
import Navbar from "./components/nav";
import { BrowserRouter } from "react-router-dom";
import PluginList from "./components/PluginList";


const queryClient = new QueryClient();

const App: React.FC = () => {
	const link = "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.RxfKUJC5hsiimFi0JhJPrgHaHa%26pid%3DApi&f=1"

  return (
    <BrowserRouter>
      <QueryClientProvider client={queryClient}>
        <Navbar />
        <PluginList data={
			[{
				"thumbnail": link,
				"name": "Surf Weather",
				"author": "mcrumb",
				"rating": "4.5",
				"description": "Get surf conditions while your coffee brews! Instaly hear about estimated wave size, wind direction, and more."
			},
			{
				"thumbnail": link,
				"name": "LocalTemps",
				"author": "jhart12354",
				"rating": "4",
				"description": "Just ask for the weather to get all the info you'll need to dress comfortably for the day."
			},
			{
				"thumbnail": link,
				"name": "Timer Clock",
				"author": "amges",
				"rating": "2.75",
				"description": "Wanna know the time? Use this plugin to ask NOVA for the time and it will give you the current time."
			}]
		} />
    </QueryClientProvider>
    </BrowserRouter>
  );
};

export default App;