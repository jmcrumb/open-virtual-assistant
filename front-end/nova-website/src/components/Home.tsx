import { Account, Profile } from '../api/accountAPI';
import { BACKEND_SRC } from '../api/helper';
import axios from 'axios';
import * as React from 'react';
import { Link } from 'react-router-dom';
import { GlobalStateContext } from '../globalState';

const defaultPhoto = "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse4.mm.bing.net%2Fth%3Fid%3DOIP.OesLvyzDO6AvU_hYUAT4IAHaHa%26pid%3DApi&f=1"

function TabLink(props) {
	const { text, url } = props



	return (
		<Link to={url}>
			<div className="TabLink">
				<span className="linkText">{text}</span>
			</div>
		</Link>
	)
}

function Home(props) {
	const [accountInfo, setAccountInfo] = React.useState({
		"name": "",
		"id": "",
	});
	const [photo, setPhoto] = React.useState([]);
	const context = React.useContext(GlobalStateContext);

	React.useEffect(() => {
		axios.get(`${BACKEND_SRC}account/${context.id}`).then((response) => {
			let acct = new Account(response.data)
			setAccountInfo({
				"name": acct.first_name,
				"id": acct.id,
			})
		});

		axios.get(`${BACKEND_SRC}acount/profile/${context.id}`).then(response => {
			let profile = new Profile(response.data)
			setPhoto(profile.photo)
		})
	}, []);

	return (
		<div className="Home">
			<div className="profileInfo">
				<img src={photo ? ("data:image/png;base64," + photo) : defaultPhoto} alt="user profile picture" className="profilePic" />
				<h1>Welcome back <span className="username">{accountInfo.name}</span>!</h1>
			</div>
			<div className="community">
				<h2>COMMUNITY</h2>
				<div className="tabs">
					<TabLink text="Publish a plugin" url="plugin/publish" />
					<TabLink text="Delete a published plugin" url="plugin/published" />
				</div>
			</div>
			<div className="myDevice">
				<h2>MY DEVICE</h2>
				<div className="tabs">
					<TabLink text="Install a plugin" url="plugin/search" />
					<TabLink text="Manage plugin settings" url="/" />
				</div>
			</div>
		</div>
	);
}

export default Home;