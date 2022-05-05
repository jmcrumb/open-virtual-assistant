import * as React from "react";
import Container from "@mui/material/Container";
import { PluginViewPublic } from "./plugin";
import PluginSearch from "./PluginSearch";

export default function Sandbox() {
  const pluginID = "e54aa9d2-0e53-471c-b9d6-f59682e5abb6";

  return (
    <Container>
      {/* <Example /> */}
      {/* <PluginViewPublic id={pluginID} /> */}
	  <PluginSearch query="t" />
    </Container>
  );
}
