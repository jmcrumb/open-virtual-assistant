import * as React from 'react';

function InputField(props, ref) {
	const {name, label, placeholder, lines, charLimit} = props;
	const input = React.useRef(null);
	const counter = React.useRef(null);

	React.useImperativeHandle(ref, () => ({
		getText: () => {
			console.log(input.current.value)
			return input.current.value
		}
	}))

	const onType = e => {
		if (!charLimit) return
		counter.current.innerText = e.target.value.length
	}

	let field;
	if (lines <= 1) {
		field = (
			<input className="field" ref={input} type="text" placeholder={placeholder} onKeyUp={onType} />
		)
	} else {
		field = (
			<textarea className="field" ref={input} name={name} id={name} rows={lines} placeholder={placeholder} onKeyUp={onType} />
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

export default React.forwardRef(InputField)