import Sandbox from "./components/sandbox";
import React from "react";
import "./styles.scss";
import "./components/PublishPlugin.css";
import "./components/PluginList.css"
import "./components/PluginPreview.css"
import "./components/Rating.scss"
import "./components/Home.css"
import { ReactQueryDevtools } from "react-query/devtools";
import Navbar from "./components/nav";
import { BrowserRouter } from "react-router-dom";
import PluginList from "./components/PluginList";

const App: React.FC = () => {
  return (
    <div className="base">
      <Navbar />
	  <Sandbox />
    </div>
  );
};

export default App;
