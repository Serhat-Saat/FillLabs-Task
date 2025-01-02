import React from "react";
import "../styles/Navbar.css";
import logo from "../assets/logo.svg";

const Navbar: React.FC = () => {
  return (
    <nav className="navbar">
      <div className="navbar-logo">
        <a
          href="https://www.fill-labs.com/"
          target="_blank"
          rel="noopener noreferrer"
        >
          <img src={logo} alt="Company Logo" className="logo" />
        </a>
        <span className="navbar-title">User Management System</span>
      </div>
    </nav>
  );
};

export default Navbar;
