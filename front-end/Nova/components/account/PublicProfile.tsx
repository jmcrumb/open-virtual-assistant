import { Component } from "react";
import { Text, View } from "react-native";
import { Account, AccountAPI, Profile } from "../../api/accountAPI";

export default class PublicProfile extends Component<{account: Account}> {

    state = {
        isLoading: true,
        profile: undefined,
        error: false
      }
    
      componentDidMount() {
        AccountAPI.getProfileByID(this.props.account.id).then((profile: Profile) => {
          this.setState({
            loading: false,
            profile: profile,
            error: false
          });
         }).catch((error) => {
          console.log(error);
          this.setState({
            loading: false,
            profile: undefined,
            error: true
           });
          });
        }
    
    render() {
        let account: Account = this.props.account;
        let profile: Profile | undefined = this.state.profile;
        return (
                <View>
                    <Text>
                        Name: {account.first_name} {account.last_name}
                    </Text>
                    {this.state.isLoading && <Text>Bio: Loading...</Text>}
                    {this.state.profile && <Text>Bio: {profile!.bio}</Text>}
                </View>
        );
    }
}