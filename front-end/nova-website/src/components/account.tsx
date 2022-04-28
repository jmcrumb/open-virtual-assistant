import axios from "axios";
import * as React from "react";
import { QueryClient, QueryClientProvider, useQuery, useQueryClient } from "react-query";
import { Account, AccountAPI } from "../api/accountAPI";
import { ReactQueryDevtools } from "react-query/devtools";
import Container from "@mui/material/Container";


const queryClient = new QueryClient();

export class AccountCard extends React.Component<{ id: string }, {account: Account}> {

    componentDidMount(): void {
        AccountAPI.getProfileByID(this.props.id)
            .then((response: any) => {
                this.setState({
                    account: response.data
                });
                console.log(response.data);
            })
            .catch((e: Error) => {
                console.log(e);
            });
    }

    render() {
        return (
            // <QueryClientProvider client={queryClient}>
            //     <p>hello</p>
            //     <AccountCardInternal id={this.props.id} />
            // </QueryClientProvider>
            <Container>
                {this.state.account.first_name}
            </Container>
        );
    }
}

// class AccountCardInternal extends React.Component<{ id: string }> {
//     render() {
//         // const { isLoading, error, data, isFetching } = useQuery("repoData", () =>
//         //     axios.get(
//         //         "https://api.github.com/repos/tannerlinsley/react-query"
//         //     ).then((res) => res.data)
//         // );

//         // if (isLoading) return "Loading...";

//         // if (error) return "An error has occurred: " + error;

//         // return (
//         //     <div>
//         //         <h1>{data.name}</h1>
//         //         <p>{data.description}</p>
//         //         <strong>👀 {data.subscribers_count}</strong>{" "}
//         //         <strong>✨ {data.stargazers_count}</strong>{" "}
//         //         <strong>🍴 {data.forks_count}</strong>
//         //         <div>{isFetching ? "Updating..." : ""}</div>
//         //         <ReactQueryDevtools initialIsOpen />
//         //     </div>
//         // );

//         // // Access the client
//         // const queryClient = useQueryClient();

//         // // Queries
//         // const { isLoading, isError, data, error } = useQuery(['account', this.props.id], () => AccountAPI.getAccountByID(this.props.id))

//         // if (isLoading) {
//         //     return <span>Loading...</span>
//         // }

//         // if (isError) {
//         //     return <span>Error</span>
//         // }

//         // // We can assume by this point that `isSuccess === true`
//         // return (
//         //     <div>{data['email']}</div>
//         // )


//     }


// }
