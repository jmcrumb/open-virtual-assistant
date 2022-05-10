import Container from "@mui/material/Container";
import { BACKEND_SRC } from "../api/helper";
import axios from "axios";
import * as React from "react";
import { useParams } from "react-router-dom";
import { Account, PublicProfile } from "../api/accountAPI";
import Avatar from "@mui/material/Avatar";
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import Typography from "@mui/material/Typography";
import Grid from "@mui/material/Grid";
import Stack from "@mui/material/Stack";
import { Plugin } from "../api/pluginStoreAPI";
import { Skeleton } from "@mui/material";

export default function ProfileView() {
    // const accountId = useParams();
    const accountId = "3bacffed-7da9-48a1-a9cf-9e89919ab0dc";

    const [profile, setProfile] = React.useState(null);

    React.useEffect(() => {
        axios.get(`${BACKEND_SRC}account/profile/${accountId}`).then((response) => {
            console.log(response.data);
            setProfile(new PublicProfile(response.data));
        });
    }, []);

    if (profile == null) return null;

    return (
        <Container>
            <Stack sx={{ paddingBottom: '2rem' }}>
                <Container>
                    <Container>
                        <ProfilePhoto img={(profile["photo"] || "")} />
                        <Typography variant="h2">{`${profile.first_name} ${profile.last_name}`}</Typography>
                    </Container>
                </Container>
                <Container>
                    <Typography>BIO</Typography>
                    <Typography>{profile.bio}</Typography>
                </Container>
                <Container>
                    <Typography>PUBLISHED PLUGINS</Typography>
                    {/* <ProfilePublishedPlugins id={profile.account_id} /> */}
                </Container>
            </Stack>

        </Container>
    );
}

function ProfilePhoto(props) {
    const img = props.img

    if (img == "") {
        return (
            <AccountCircleIcon sx={{ width: 56, height: 56 }} />
        );
    } else {
        return (
            <Avatar
                alt="Profile Picture"
                src={img}
                sx={{ width: 56, height: 56 }}
            />
        );
    }
}

function ProfilePublishedPlugins(props) {
    const accountId: string = props.id;

    const [plugins, setPlugins] = React.useState(null);
    
    React.useEffect(() => {
        axios.get(`${BACKEND_SRC}plugin/search/account/${accountId}`).then((response) => {
            let instantiated_plugins = [];

            response.data.forEach((value) => {
                instantiated_plugins.push(new Plugin(value));
            });

            setPlugins(instantiated_plugins);
        });
    });

    if(!plugins) return <Skeleton />;

    return (
        <Stack>
            {plugins.map((plugin) => { (<Typography>{plugin.name}</Typography>) })}
        </Stack>
    );
}