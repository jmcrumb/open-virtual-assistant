import { AccountAPI, useQueryAccountByID } from "../api/accountAPI";
import * as React from "react";
import { AccountCard } from "./account";
import Container from "@mui/material/Container";
import { PluginViewPublic } from "./plugin";
import Login from "./login";

export default function Sandbox() {
  const pluginID = "3f094753-6d45-4897-a749-c51378ddbe13";

  return (
    <Container>
      {/* <Example /> */}
      <Login />
    </Container>
  );
}
