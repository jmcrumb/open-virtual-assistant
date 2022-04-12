import * as React from "react";
import { AccountCard } from "./account";

let Logo = "https://logrocket-assets.io/static/home-hero-c97849b227a3d3015730e3371a76a7f0.svg";
export default class Sandbox extends React.Component<{}> {
    render() {

        let accountId: string = "425bb481-dd67-41e5-820a-5ec0c6cc2bbd";

        return (
            <div>
                <h3>A Simple React Component Example with Typescript</h3>
                <div>
                    <img height="250" src={Logo} />
                </div>
                <p>This component shows the Logrocket logo.</p>
                <p>For more info on Logrocket, please visit https://logrocket.com </p>
                <br />
                <AccountCard id={accountId}/>
            </div>
        );
    }
}
