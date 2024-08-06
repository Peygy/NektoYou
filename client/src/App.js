import React, { useState } from 'react';
import { gql, useMutation } from '@apollo/client';

const REGISTER_USER_MUTATION = gql`
  mutation RegisterUser($input: UserInput!) {
    registerUser(input: $input) {
      accessToken
      refreshToken
    }
  }
`;

function App() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [registerUser, { data, loading, error }] = useMutation(REGISTER_USER_MUTATION);

  const handleSubmit = (e) => {
    e.preventDefault();
    registerUser({ variables: { input: { username, password } } });
  };

  return (
    <div>
      <h1>GraphQL Client</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={username}
          placeholder="Username"
          onChange={(e) => setUsername(e.target.value)}
        />
        <input
          type="text"
          value={password}
          placeholder="Password"
          onChange={(e) => setPassword(e.target.value)}
        />
        <button type="submit">Submit</button>
      </form>
      {loading && <p>Loading...</p>}
      {error && <p>Error: {error.message}</p>}
      {data && (
        <div>
          <p>Access Token: {data.registerUser.accessToken}</p>
          <p>Refresh Token: {data.registerUser.refreshToken}</p>
        </div>
      )}
    </div>
  );
}

export default App;
