import { Plugin } from "../api/pluginStoreAPI"
import * as React from "react";
import axios from "axios";
import { BACKEND_SRC } from "../api/helper";
import PluginList from "./PluginList";
import InputField from "./InputField";

const link = "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.RxfKUJC5hsiimFi0JhJPrgHaHa%26pid%3DApi&f=1"

function PluginSearch() {
	const [data, setData] = React.useState([])
	const search = React.useRef(null)

	const onSearch = () => {
		axios.get(`${BACKEND_SRC}plugin/search/${search.current.getText()}`)
		.then(response => {
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
			setData(temp);
		})
	}

	return (
		<div className="PluginSearch">
			<div className="search">
				<InputField ref={search} name="search-bar" label="Search plugins" lines={1} placeholder="plugin name" />
				<button className="submit" onClick={onSearch}>Search</button>
			</div>
			<PluginList data={data} />
		</div>
	);
}

export default PluginSearch;