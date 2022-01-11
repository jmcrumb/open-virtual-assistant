import { createContext } from 'react';
import NovaCore from '../nova-core/core';

const CoreContext = createContext(new NovaCore());
export default CoreContext;