import * as React from 'react';
import { StyleSheet } from 'react-native';
import { Account, AccountAPI } from '../api/accountAPI';
import PublicProfile from '../components/account/PublicProfile';

import EditScreenInfo from '../components/EditScreenInfo';
import { Text, View } from '../components/Themed';

export default class AccountViewScreen extends React.Component {

  state = {
    isLoading: true,
      account: null,
      error: false
  }

  componentDidMount() {
    AccountAPI.getAccountByID('01d2c9e4-69a8-40e0-b2ba-25f78822c6dd').then((account: Account) => {
      this.setState({
        loading: false,
        account: account,
        error: false
      });
      console.log(account);
     }).catch((error) => {
      console.log(error);
      this.setState({
        loading: false,
        account: null,
        error: true
       });
      });
  }

  render() {
    return (
      <View style={styles.container}>
        <Text style={styles.title}>My Account</Text>
        <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
        <EditScreenInfo path="/screens/AccountViewScreen.tsx" />
        {this.state.isLoading && <Text>Loading...</Text>}
        {this.state.account && <PublicProfile account={this.state.account}></PublicProfile>}
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
