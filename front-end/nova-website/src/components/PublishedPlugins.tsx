import { Plugin } from "../api/pluginStoreAPI";
import * as React from "react";
import axios from "axios";
import { BACKEND_SRC } from "../api/helper";
import PluginList from "./PluginList";
import { GlobalStateContext } from "../globalState";

const link = "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.RxfKUJC5hsiimFi0JhJPrgHaHa%26pid%3DApi&f=1"

function PublishedPlugins() {
	const [plugins, setPlugins] = React.useState([])
	const context = React.useContext(GlobalStateContext);

	React.useEffect(() => {
		axios.get(`${BACKEND_SRC}plugin/publishedBy/${context.id}`).then((response) => {
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

	return (
		<div className="PublishedPlugins">
			<PluginList data={plugins} />
		</div>
	);
}

export default PublishedPlugins;