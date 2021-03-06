import React from "react";

const defaultGlobalState = {
    id: "",
    setId: () => {},
    token: "",
    setToken: () => {}
};

interface GlobalStateInterface {
    id: string,
    setId?,
    token: string,
    setToken?
};

const GlobalStateContext = React.createContext<Partial<GlobalStateInterface>>({});

const GlobalStateProvider = ({ children }) => {
    const [id, setIdInternal] = React.useState("86ee5cd6-5c83-4fd3-b4d6-1c2064dcd918");
    const [token, setTokenInternal] = React.useState("");

    const setId = ((val) => { setIdInternal(val) });
    const setToken = ((val) => { setTokenInternal(val) });

    return (
        <GlobalStateContext.Provider value={{
            id: id,
            setId: setId,
            token: token,
            setToken: setToken
        }}>
            {children}
        </GlobalStateContext.Provider>
    );
};

export { GlobalStateProvider, GlobalStateContext }

// const globalStateContext = React.createContext(defaultGlobalState);
// const dispatchStateContext = React.createContext(undefined);

// const GlobalStateProvider = ({ children }) => {
//     const [state, dispatch] = React.useReducer(
//         (state, newValue) => ({ ...state, ...newValue }),
//         defaultGlobalState
//     );
//     return (
//         <globalStateContext.Provider value={state}>
//             <dispatchStateContext.Provider value={dispatch}>
//                 {children}
//             </dispatchStateContext.Provider>
//         </globalStateContext.Provider>
//     );
// };

// const useGlobalState = () => [
//     React.useContext(globalStateContext),
//     React.useContext(dispatchStateContext)
// ];

// export {GlobalStateProvider, useGlobalState}