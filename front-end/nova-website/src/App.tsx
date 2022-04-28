import Sandbox from "./components/sandbox";
import React from "react";
import "./styles.scss";
import { ReactQueryDevtools } from "react-query/devtools";
import { QueryClient, QueryClientProvider} from "react-query";
import Navbar from "./components/nav";
import { BrowserRouter } from "react-router-dom";


const queryClient = new QueryClient();

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <QueryClientProvider client={queryClient}>
        <Navbar />
        <Sandbox />
        <ReactQueryDevtools initialIsOpen />
    </QueryClientProvider>
    </BrowserRouter>
  );
};

export default App;