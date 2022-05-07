import { Plugin, Review } from "../api/pluginStoreAPI";
import * as React from "react";
import axios from "axios";
import { BACKEND_SRC } from "../api/helper";
import PluginList from "./PluginList";

const link = "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Ftse1.mm.bing.net%2Fth%3Fid%3DOIP.RxfKUJC5hsiimFi0JhJPrgHaHa%26pid%3DApi&f=1"

function InstallePlugins(props) {
	const [plugins, setPlugins] = React.useState([])

	React.useEffect(() => {	
		// query installed plugins from local webserver
	  }, []);

	return (
		<div className="InstallePlugins">
			<PluginList data={plugins} />
		</div>
	);
}

export default InstallePlugins;