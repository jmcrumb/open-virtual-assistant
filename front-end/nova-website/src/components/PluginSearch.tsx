import { Plugin, Review } from "../api/pluginStoreAPI";
import * as React from "react";
import axios from "axios";
import { BACKEND_SRC } from "../api/helper";
import PluginList from "./PluginList";
import { useParams } from "react-router-dom";

function PluginSearch() {
	const { query } = useParams();

	return (
		<div className="PluginSearch">
			<PluginList data={ query } />
		</div>
	);
}

export default PluginSearch;