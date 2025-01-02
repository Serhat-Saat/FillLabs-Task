import React from "react";
import Navbar from "./components/Navbar";
import UserList from "./components/UserList";
import "./styles/App.css";

const App: React.FC = () => {
  return (
    <div>
      <Navbar />
      <div className="container">
        <UserList />
      </div>
    </div>
  );
};

export default App;
