import { AccountAPI, useQueryAccountByID } from "../api/accountAPI";
import * as React from "react";
import ProfileView from "./profile";
import Container from "@mui/material/Container";

export default function Sandbox() {
  const pluginID = "3f094753-6d45-4897-a749-c51378ddbe13";

  return (
    <Container>
      <ProfileView />
    </Container>
  );
}
