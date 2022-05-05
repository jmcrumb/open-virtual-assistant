import * as React from 'react';
import PluginPreview from './PluginPreview';

function PluginList(props) {
	let {data} = (props) ? props : []
	let plugins = []

	for (let plugin of data) {
		plugins.push(
			<PluginPreview id={plugin.id} thumbnail={plugin.thumbnail} name={plugin.name} author={plugin.author} rating={plugin.rating} description={plugin.description}/>
		)
	}

	return (
		<div className="PluginList">
			{plugins}
		</div>
	);
}

export default PluginList;