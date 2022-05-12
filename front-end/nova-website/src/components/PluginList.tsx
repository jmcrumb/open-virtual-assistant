import { BACKEND_SRC } from '../api/helper';
import * as React from 'react';
import PluginPreview from './PluginPreview';
import axios from 'axios';
import { Plugin } from "../api/pluginStoreAPI";

const link = "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.RxfKUJC5hsiimFi0JhJPrgHaHa%26pid%3DApi&f=1"

function PluginList(props) {
	let { data } = props
	const [plugins, setPlugins] = React.useState([])

	React.useEffect(() => {	
		axios.get(`${BACKEND_SRC}plugin/search/${data.toString()}`).then((response) => {
			if (!response.data) {
				return
			}
			let temp = []
			response.data.forEach(p => {
				p = new Plugin(p)
				temp.push({
					"id": p.id,
					"thumbnail": link,
					"name": p.name,
					"author": p.publisher,
					"rating": 5,
					"description": p.about,
				});
			});
			setPlugins(temp);
		});
	  }, []);

	for (let plugin of data) {
		plugins.push(
			<PluginPreview
				id={plugin.id}
				thumbnail={plugin.thumbnail}
				name={plugin.name}
				author={plugin.author}
				rating={plugin.rating}
				description={plugin.description}
			/>
		)
	}

	return (
		<div className="PluginList">
			{plugins}
		</div>
	);
}

export default PluginList;