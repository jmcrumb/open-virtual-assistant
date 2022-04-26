import Sandbox from "./components/sandbox";
import React from "react";
import "./styles.scss";
import { ReactQueryDevtools } from "react-query/devtools";
import { QueryClient, QueryClientProvider} from "react-query";


const queryClient = new QueryClient();

const App: React.FC = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <Sandbox />
      <ReactQueryDevtools initialIsOpen />
    </QueryClientProvider>
  );
};

export default App;