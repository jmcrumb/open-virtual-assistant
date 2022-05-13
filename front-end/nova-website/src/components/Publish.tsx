import { BACKEND_SRC } from "../api/helper";
import { Plugin } from "../api/pluginStoreAPI";
import axios from "axios";
import * as React from "react";
import { useNavigate } from "react-router-dom";
import { GlobalStateContext } from "../globalState";
import InputField from "./InputField";

function Publish() {
	const context = React.useContext(GlobalStateContext);
	const sourceLink = React.useRef(null);
	const pluginName = React.useRef(null);
	const description = React.useRef(null);
	let navigate = useNavigate();

	const onPublish = () => {
		let plugin = new Plugin({
			"publisher": context.id,
			"name": pluginName.current.getText(),
			"source_link": sourceLink.current.getText(),
			"about": description.current.getText(),
		})
		console.log(JSON.stringify(plugin))

		axios.post(`${BACKEND_SRC}plugin/`, plugin)
		.then(() => {
			navigate(`/plugin/published`, {replace: true})
		})
		.catch(() => {
			navigate("/", {replace: true})
		})
	}

	return (
		<div className="Publish">
			<h1>Publish a new plugin</h1>
			<InputField ref={sourceLink} name="source-link" label="link source code" lines={1} placeholder="source code link" />
			<InputField ref={pluginName} name="plugin-name" label="name your plugin" lines={1} placeholder="plugin name" />
			<InputField ref={description} name="description" label="give it a description" lines={5} charLimit={500} placeholder="description" />
			<button className="publish" onClick={onPublish}>publish</button>
		</div>
	)
}

export default Publish;