import { AccountAPI, useQueryAccountByID } from "../api/accountAPI";
import * as React from "react";
import { useQueryClient } from "react-query";
import { AccountCard } from "./account";
import Container from "@mui/material/Container";
import { PluginViewPublic } from "./plugin";

export default function Sandbox() {
  const pluginID = "3f094753-6d45-4897-a749-c51378ddbe13";

  return (
    <Container>
      {/* <Example /> */}
      <PluginViewPublic id={pluginID} />
    </Container>
  );
}

function Example() {
  const queryClient = useQueryClient();
  const { status, data, error, isFetching } = useQueryAccountByID();

  return (
    <div>
      <div>
        {status === "loading" ? (
          "Loading..."
        ) : status === "error" ? (
          <span>Error: {error}</span>
        ) : (
          <>
            <div>
              <p key={data.id}>
                <a
                  onClick={() => alert("action")}
                  href="#"
                  style={
                    // We can access the query data here to show bold links for
                    // ones that are cached
                    queryClient.getQueryData(["post", data.id])
                      ? {
                          fontWeight: "bold",
                          color: "green",
                        }
                      : {}
                  }
                >
                  {data.email}
                </a>
              </p>
            </div>
            <div>{isFetching ? "Background Updating..." : " "}</div>
          </>
        )}
      </div>
    </div>
  );
}
