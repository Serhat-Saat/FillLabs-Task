import React, { useEffect, useState } from "react";
import { FaEdit, FaPlus, FaTrash } from "react-icons/fa";
import UserModal from "./UserModal";
import "../styles/UserList.css";

interface User {
  id: number;
  userName: string;
  userEmail: string;
  userPhone: string;
}

const UserList: React.FC = () => {
  const [users, setUsers] = useState<User[]>([]);
  const [selectedUser, setSelectedUser] = useState<User | null>(null);
  const [modalAction, setModalAction] = useState<
    "create" | "edit" | "delete" | null
  >(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchUsers(); // Fetch users when the component mounts
  }, []);
  // Function to fetch users from the API
  const fetchUsers = async () => {
    try {
      const response = await fetch("http://localhost:8080/users");
      if (!response.ok) throw new Error("Failed to fetch users.");
      const data = await response.json();
      setUsers(data);
    } catch {
      setError("Error fetching users.");
    }
  };
  // Handle user deletion by removing them from the state
  const handleDelete = (id: number) => {
    setUsers((prevUsers) => prevUsers.filter((user) => user.id !== id));
    setModalAction(null);
  };
  // Function to format phone numbers into the desired format
  const formatPhoneNumber = (phone: string) => {
    const numericValue = phone.replace(/\D/g, "");
    if (numericValue.length <= 3) return `(${numericValue})`;
    return `(${numericValue.slice(0, 3)})${numericValue.slice(3, 10)}`;
  };
  // Function to close the modal
  const handleCloseModal = () => setModalAction(null);

  return (
    <div className="user-list-container">
      {error && <div className="error">{error}</div>}
      <div className="actions">
        <button onClick={() => setModalAction("create")}>
          <FaPlus /> New
        </button>
        <button
          onClick={() => selectedUser && setModalAction("edit")}
          disabled={!selectedUser || modalAction !== null}
        >
          <FaEdit /> Edit
        </button>
        <button
          onClick={() => selectedUser && setModalAction("delete")}
          disabled={!selectedUser || modalAction !== null}
        >
          <FaTrash /> Delete
        </button>
      </div>
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Email</th>
            <th>Phone</th>
          </tr>
        </thead>
        <tbody>
          {users.map((user) => (
            <tr
              key={user.id}
              onClick={() => setSelectedUser(user)}
              className={selectedUser?.id === user.id ? "selected" : ""}
            >
              <td>{user.userName}</td>
              <td>{user.userEmail}</td>
              <td>{formatPhoneNumber(user.userPhone)}</td>
            </tr>
          ))}
        </tbody>
      </table>
      {modalAction && (
        <UserModal
          setSelectedUser={setSelectedUser}
          user={modalAction === "create" ? null : selectedUser} // Pass selected user to the modal (if editing)
          action={modalAction}
          onClose={handleCloseModal}
          onUserUpdated={(updatedUser) => {
            if (modalAction === "create") {
              setUsers((prevUsers) => [...prevUsers, updatedUser]);
            } else {
              setUsers((prevUsers) =>
                prevUsers.map((user) =>
                  user.id === updatedUser.id ? updatedUser : user
                )
              );
            }
            handleCloseModal();
          }}
          onUserDeleted={handleDelete}
        />
      )}
    </div>
  );
};

export default UserList;
