import React from 'react';
import './App.css';
import FindHighestPrime from './number/FindHighestPrime';

export function App() {
  return (
    <div className="app">
      <h2 className="title">Welcome to number-app</h2>
      <div>
        <FindHighestPrime />
      </div>
    </div>
  );
}
