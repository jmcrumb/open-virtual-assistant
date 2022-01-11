import * as React from 'react';
import { Button, StyleSheet } from 'react-native';

import EditScreenInfo from '../components/EditScreenInfo';
import { Text, View } from '../components/Themed';
import CoreContext from '../constants/context';
import { RootTabScreenProps } from '../types';

export default function AssistantViewScreen({ navigation }: RootTabScreenProps<'AssistantView'>) {

  return (
    <View style={styles.container}>
      <Text style={styles.title}>NOVA</Text>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <EditScreenInfo path="/screens/AssistantViewScreen.tsx" />
      <AssistantInteractor /> 
    </View>
  );
}

class AssistantInteractor extends React.Component {

  state = {response: ''};

  render(): React.ReactNode {
      return (
        <View>
          <Button 
            title="Say Hi"
            onPress={() => {
              const core = React.useContext(CoreContext);
              this.state.response = core.invoke('Hello');
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
});
