import { BACKEND_SRC } from "../api/helper";
import { Review, Plugin } from "../api/pluginStoreAPI";
import axios from "axios";
import * as React from "react";
import UserState from "../userState"
import { getTextFieldUtilityClass } from "@mui/material";
import { useNavigate } from "react-router-dom";

function inputField(props, ref) {
	const {name, label, placeholder, lines, charLimit} = props
	const input = React.useRef(null)
	const counter = React.useRef(null)

	React.useImperativeHandle(ref, () => ({
		getText: () => {
			console.log(input.current.value)
			return input.current.innerText
		}
	}))

	const onType = e => {
		if (!charLimit) return
		counter.current.innerText = e.target.value.length
	}

	React.useEffect(() => {
		input.current.addEventListener("keyup", onType)
	}, [])

	let field;
	if (lines <= 1) {
		field = (
			<input className="field" ref={input} type="text" placeholder={placeholder} />
		)
	} else {
		field = (
			<textarea className="field" ref={input} name={name} id={name} rows={lines} placeholder={placeholder} />
		)
	}

	let charCounter = (
		<div className="charCount">
			<span ref={counter}>{counter.current ? counter.current : 0}</span>/{charLimit} characters
		</div>
	)

	return (
		<div className="InputField">
			<h2>{label}</h2>
			{field}
			{charLimit ? charCounter : ""}
		</div>
	)
}
const InputField = React.forwardRef(inputField)

function PublishPlugin() {
	const sourceLink = React.useRef(null);
	const pluginName = React.useRef(null);
	const description = React.useRef(null);
	let navigate = useNavigate();

	const onPublish = () => {
		let plugin = new Plugin({
			"publisher": UserState.getInstance().state["id"],
			"name": sourceLink.current.getText(),
			"sourceLink": pluginName.current.getText(),
			"about": description.current.getText(),
		})
		axios.post(`${BACKEND_SRC}plugin`, plugin)
		// need error handling and feedback to user about whether
		// the request went through or not
		navigate("/plugin/published", {replace: true});
	}

	return (
		<div className="PublishPlugin">
			<h1>Publish a new plugin</h1>
			<InputField ref={sourceLink} name="source-link" label="link source code" lines={1} placeholder="source code link" />
			<InputField ref={pluginName} name="plugin-name" label="name your plugin" lines={1} placeholder="plugin name" />
			<InputField ref={description} name="description" label="give it a description" lines={5} charLimit={500} placeholder="description" />
			<button className="publish" onClick={onPublish}>publish</button>
		</div>
	)
}

export default PublishPlugin;