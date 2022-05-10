import { BACKEND_SRC } from "../api/helper";
import { Review, Plugin } from "../api/pluginStoreAPI";
import axios from "axios";
import * as React from "react";
import UserState from "../userState"

function InputField(props) {
	const {name, label, placeholder, lines, charLimit} = props
	const input = React.useRef(null)
	const counter = React.useRef(null)

	const onType = e => {
		console.log("burger")
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

function PublishPlugin() {
	const sourceLink = React.useRef(null)
	const pluginName = React.useRef(null)
	const description = React.useRef(null)

	const getInputText = input => {
		return input.current.querySelector(".field").value
	}
	const onPublish = () => {
		let plugin = new Plugin({
			"publisher": UserState.getInstance().state["id"],
			"name": getInputText(sourceLink),
			"sourceLink": getInputText(pluginName),
			"about": getInputText(description),
		})
		axios.post(`${BACKEND_SRC}plugin`, plugin)
		// need error handling and feedback to user about whether
		// the request went through or not
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