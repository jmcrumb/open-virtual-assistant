import * as React from "react";
import Container from "@mui/material/Container";
import { PluginViewPublic } from "./plugin";
import PluginSearch from "./PluginSearch";
import PublishedPlugins from "./PublishedPlugins";
import Login from "./Login";
import Home from "./Home";
import { AccountCard } from "./account";
import Container from "@mui/material/Container";
import { PluginViewPublic } from "./plugin";
import Login from "./login";

export default function Sandbox() {

  return (
    <Container>
      {/* <Example /> */}
      {/* <PluginViewPublic id={pluginID} /> */}
	  {/* <PublishedPlugins query="86ee5cd6-5c83-4fd3-b4d6-1c2064dcd918" /> */}
	  {/* <Login /> */}
	  <Home />
    </Container>
  );
}
