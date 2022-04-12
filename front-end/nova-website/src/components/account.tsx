import * as React from "react";
import { QueryClient, QueryClientProvider, useQuery, useQueryClient } from "react-query";
import { Account, AccountAPI } from "../api/accountAPI";

const queryClient = new QueryClient();

export class AccountCard extends React.Component<{ id: string }> {
    render() {
        return (
            <QueryClientProvider client={queryClient}>
                <AccountCardInternal id={this.props.id} />
            </QueryClientProvider>
        );
    }
}

class AccountCardInternal extends React.Component<{ id: string }> {
    render() {
        // Access the client
        const queryClient = useQueryClient();

        // Queries
        const { isLoading, isError, data, error } = useQuery(['account', this.props.id], () => AccountAPI.getAccountByID(this.props.id))

        if (isLoading) {
            return <span>Loading...</span>
        }

        if (isError) {
            return <span>Error</span>
        }

        // We can assume by this point that `isSuccess === true`
        return (
            <ul>
                {data.map((account: Account) => (
                    <li key={account.id}>{account.email}</li>
                ))}
            </ul>
        )
    }


}
