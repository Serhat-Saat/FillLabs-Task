import React, { Dispatch, SetStateAction, useState } from "react";
import "../styles/UserModal.css";

interface User {
  id: number;
  userName: string;
  userEmail: string;
  userPhone: string;
}

interface UserModalProps {
  user: User | null;
  action: "create" | "edit" | "delete";
  setSelectedUser: Dispatch<SetStateAction<User | null>>;
  onClose: () => void;
  onUserUpdated: (user: User) => void;
  onUserDeleted: (id: number) => void;
}

const UserModal: React.FC<UserModalProps> = ({
  user,
  action,
  setSelectedUser,
  onClose,
  onUserUpdated,
  onUserDeleted,
}) => {
  const [userName, setUserName] = useState(user?.userName || "");
  const [userEmail, setUserEmail] = useState(user?.userEmail || "");
  const [userPhone, setUserPhone] = useState(user?.userPhone || "");
  const [showDeleteConfirmation, setShowDeleteConfirmation] = useState(false);
  const [isProcessing, setIsProcessing] = useState(false);
  const [phoneError, setPhoneError] = useState<string | null>(null);

  // Helper function to format phone numbers
  const formatPhoneNumber = (input: string) => {
    const numericValue = input.replace(/\D/g, "");
    if (numericValue.length === 0) return "(5)";
    if (numericValue.length <= 3) return `(${numericValue.slice(0, 3)})`;
    return `(${numericValue.slice(0, 3)})${numericValue.slice(3, 10)}`;
  };

  // Handles phone input change and validation
  const handlePhoneChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const formattedPhone = formatPhoneNumber(e.target.value);
    setUserPhone(formattedPhone);

    const numericPhone = formattedPhone.replace(/\D/g, "");
    if (numericPhone.length < 10) {
      setPhoneError("Phone number must be 10 digits!");
    } else {
      setPhoneError(null);
    }
  };

  // Handles form submission for creating, editing, or deleting users
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    // Prevent form submission if the request is already being processed
    if (isProcessing) return;
    setIsProcessing(true);

    // If action is delete, show delete confirmation modal
    if (action === "delete") {
      setShowDeleteConfirmation(true);
      setIsProcessing(false);
      return;
    }

    const numericPhone = userPhone.replace(/\D/g, "");
    // Validate phone number length before submitting
    if (numericPhone.length !== 10) {
      setPhoneError("Phone number must be 10 digits");
      setIsProcessing(false);
      return;
    } else {
      setPhoneError(null);
    }

    const newUser = {
      id: user?.id || Date.now(), // Use existing user ID or generate a new one
      userName,
      userEmail,
      userPhone: userPhone.replace(/\D/g, ""),
    };

    try {
      // Determine URL and HTTP method based on action (create or update)
      const url =
        action === "create"
          ? "http://localhost:8080/users/create"
          : "http://localhost:8080/users/update";
      const method = action === "create" ? "POST" : "PUT";
      const response = await fetch(url, {
        method,
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(newUser),
      });
      if (!response.ok) throw new Error("Failed to save user.");
      const updatedUser = await response.json();
      onUserUpdated(updatedUser);
    } catch (error) {
      console.error("Error saving user:", error);
    } finally {
      setIsProcessing(false);
    }
  };
  // Handles user deletion confirmation
  const handleDeleteConfirm = async () => {
    if (isProcessing || !user) return;
    setIsProcessing(true);
    try {
      const response = await fetch(
        `http://localhost:8080/users/delete?id=${user.id}`,
        {
          method: "DELETE",
          headers: { "Content-Type": "application/json" },
        }
      );
      if (!response.ok) throw new Error("Failed to delete user.");
      onUserDeleted(user.id);
      setSelectedUser(null);
      setShowDeleteConfirmation(false);
      onClose();
    } catch (error) {
      console.error("Error deleting user:", error);
    } finally {
      setIsProcessing(false);
    }
  };

  return (
    <div className="modal-overlay">
      <div className="modal-content">
        {!showDeleteConfirmation ? (
          <>
            <h2>
              {action === "create"
                ? "Create User"
                : action === "edit"
                ? "Edit User"
                : "Delete User"}
            </h2>
            <form onSubmit={handleSubmit}>
              <label>
                Name:
                <input
                  placeholder="John Doe"
                  type="text"
                  value={userName}
                  onChange={(e) => setUserName(e.target.value)}
                  required
                  readOnly={action === "delete"}
                />
              </label>
              <label>
                Email:
                <input
                  placeholder="johndoe@example.com"
                  type="email"
                  value={userEmail}
                  onChange={(e) => setUserEmail(e.target.value)}
                  required
                  readOnly={action === "delete"}
                />
              </label>
              <label>
                Phone:
                <input
                  placeholder="(5xx)xxxxxxx"
                  type="text"
                  value={userPhone}
                  onChange={handlePhoneChange}
                  required
                  readOnly={action === "delete"}
                  maxLength={14}
                />
                {phoneError && <p className="error-message">{phoneError}</p>}
              </label>
              <div className="modal-actions">
                {action === "delete" ? (
                  <>
                    <button
                      type="submit"
                      className="delete-button"
                      disabled={isProcessing}
                    >
                      Delete
                    </button>
                    <button
                      type="button"
                      onClick={onClose}
                      className="cancel-button"
                    >
                      Back
                    </button>
                  </>
                ) : (
                  <>
                    <button
                      type="submit"
                      className={
                        action === "edit" ? "save-button" : "create-button"
                      }
                      disabled={isProcessing}
                    >
                      {action === "create" ? "Create" : "Save"}
                    </button>
                    <button
                      type="button"
                      onClick={onClose}
                      className="cancel-button"
                    >
                      Back
                    </button>
                  </>
                )}
              </div>
            </form>
          </>
        ) : (
          <div className="confirmation-modal-overlay">
            <div className="confirmation-modal-content">
              <p>Are you sure you want to delete this user?</p>
              <div className="confirmation-modal-actions">
                <button
                  onClick={handleDeleteConfirm}
                  className="delete-button"
                  disabled={isProcessing}
                >
                  Yes
                </button>
                <button
                  onClick={() => setShowDeleteConfirmation(false)}
                  className="cancel-button"
                >
                  Back
                </button>
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default UserModal;
