import * as React from 'react';
import { Button, StyleSheet, TextInput } from 'react-native';

import EditScreenInfo from '../components/EditScreenInfo';
import { Text, View } from '../components/Themed';
import CoreContext from '../constants/context';
import NovaCore from '../nova-core/core';
import { RootTabScreenProps } from '../types';

export default function AssistantViewScreen({ navigation }: RootTabScreenProps<'AssistantView'>) {

  return (
    <View style={styles.container}>
      <Text style={styles.title}>NOVA</Text>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <AssistantInteractor /> 
    </View>
  );
}

class AssistantInteractor extends React.Component {
  state: {[key: string]: any};
  static contextType = CoreContext;

  constructor(props: any) {
    super(props);
    this.state = {command: '', response: ''};
  }

  handleCommand = (text: string) => {
    this.setState({'command': text});
  }

  render(): React.ReactNode {
      return (
        <View>
          <TextInput
            style={styles.input}
            onChangeText={this.handleCommand}
            placeholder="Ask me something"
          />
          <Button 
            title="Say Hi"
            onPress={() => {
              this.setState({'response': this.context.invoke(this.state.command)});
            }} 
          />
          <Text>{this.state.response}</Text>
        </View>
      );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
  },
  title: {
    fontSize: 20,
    fontWeight: 'bold',
  },
  separator: {
    marginVertical: 30,
    height: 1,
    width: '80%',
  },
  input: {
    height: 40,
    margin: 12,
    borderWidth: 1,
    padding: 10,
  },
});
