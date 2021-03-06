import * as React from 'react';
import PluginPreview from './PluginPreview';

function PluginList(props) {
	const {data} = props

	// Put this in search plugins
	// const link = "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.RxfKUJC5hsiimFi0JhJPrgHaHa%26pid%3DApi&f=1"
	// React.useEffect(() => {
	// 	const data = new URLSearchParams(useLocation().search).get("query");
	// 	axios.get(`${BACKEND_SRC}plugin/search/${JSON.stringify(data)}`).then((response) => {
	// 		if (!response.data) {
	// 			return
	// 		}
	// 		let temp = []
	// 		response.data.forEach(p => {
	// 			p = new Plugin(p)
	// 			temp.push({
	// 				"id": p.id,
	// 				"thumbnail": link,
	// 				"name": p.name,
	// 				"author": p.publisher,
	// 				"rating": 5,
	// 				"description": p.about,
	// 			});
	// 		});
	// 		setPlugins(temp);
	// 	});
	//   }, []);

	let pluginElements = []
	for (let plugin of data) {
		pluginElements.push(
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
			{
			(pluginElements) ?
			(
				pluginElements
			) :
			(
				<div className="placeholder">
					Search for a plugin in the search bar!
				</div>
			)
			}
		</div>
	);
}

export default PluginList;