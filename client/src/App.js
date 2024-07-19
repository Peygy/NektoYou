import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

const Home = () => <h1>Hello home</h1>;
const Main = () => <h1>Hello world</h1>;

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Main />} />
        <Route path="/home" element={<Home />} />
      </Routes>
    </Router>
  );
}

export default App;