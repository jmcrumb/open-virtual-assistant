import { Plugin, Review } from "../api/pluginStoreAPI";
import * as React from "react";
import axios from "axios";
import { BACKEND_SRC } from "../api/helper";
import PluginList from "./PluginList";

const link = "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.RxfKUJC5hsiimFi0JhJPrgHaHa%26pid%3DApi&f=1"

function PluginSearch(props) {
	const [plugins, setPlugins] = React.useState([])

	React.useEffect(() => {	
		axios.get(`${BACKEND_SRC}plugin/search/${props.query}`).then((response) => {
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
		<div className="PluginSearch">
			<PluginList data={plugins} />
		</div>
	);
}

export default PluginSearch;