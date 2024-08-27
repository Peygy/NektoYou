import React from 'react';
import ReactDOM from 'react-dom';
import { ApolloClient, InMemoryCache, ApolloProvider } from '@apollo/client';
import App from './App';

const authClient = new ApolloClient({
  uri: 'http://localhost:4000/graphql/register',
  cache: new InMemoryCache()
});

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <ApolloProvider client={authClient}>
    <App />
  </ApolloProvider>
);