import { Container } from "@material-ui/core";
import { AccountAPI } from "../api/accountAPI";
import * as React from "react";
import { useQueryClient } from "react-query";
import { AccountCard } from "./account";

let Logo = "https://logrocket-assets.io/static/home-hero-c97849b227a3d3015730e3371a76a7f0.svg";
export default class Sandbox extends React.Component<{}> {
    render() {

        let accountId: string = "425bb481-dd67-41e5-820a-5ec0c6cc2bbd";

        return (
            <Container>
                <p>Account Test</p>
                <Example />
            </Container>
        );
    }
}

function Example() {

    const queryClient = useQueryClient();
    const { status, data, error, isFetching } = AccountAPI.useQueryAccountByID();
  
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
