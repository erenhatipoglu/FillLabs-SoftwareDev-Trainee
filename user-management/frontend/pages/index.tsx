import React, { useEffect, useState } from "react"; // Import React and hooks
import axios from "axios"; // Import axios for API requests
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Button,
  TextField,
  Box,
} from "@mui/material"; // Import Material-UI components for styling

// Define the User interface to type-check user objects
interface User {
  id: number; // User ID
  name: string; // User's name
  email: string; // User's email
}

const Home: React.FC = () => {
  // State to manage the list of users
  const [users, setUsers] = useState<User[]>([]);
  // State to manage the currently selected user for editing
  const [selectedUser, setSelectedUser] = useState<User | null>(null);
  // State to manage the form inputs for name and email
  const [form, setForm] = useState({ name: "", email: "" });

  // useEffect to fetch users when the component mounts
  useEffect(() => {
    fetchUsers();
  }, []);

  // Fetch users from the API
  const fetchUsers = async () => {
    const response = await axios.get<User[]>("/api/users"); // GET request to fetch users
    setUsers(response.data); // Update the users state with the response data
  };

  // Handle saving a new or edited user
  const handleSave = async () => {
    // Validate that both name and email are provided
    if (!form.name || !form.email) {
      alert("Name and Email are required");
      return;
    }

    // Check for duplicate email in the current user list
    const existingUser = users.find((user) => user.email === form.email);
    if (existingUser && (!selectedUser || existingUser.id !== selectedUser.id)) {
      alert("User with this email already exists");
      return;
    }

    try {
      // Update user if one is selected, otherwise create a new user
      if (selectedUser) {
        await axios.put(`/api/users/${selectedUser.id}`, form); // PUT request for updating user
      } else {
        await axios.post("/api/users", form); // POST request for creating new user
      }
      fetchUsers(); // Refresh user list after the save
      setForm({ name: "", email: "" }); // Clear the form
      setSelectedUser(null); // Reset selected user
    } catch (error) {
      console.error("Error saving user:", error); // Log error to console
      alert("An error occurred while saving the user."); // Show error message
    }
  };

  // Handle editing a user
  const handleEdit = (user: User) => {
    setSelectedUser(user); // Set the selected user for editing
    setForm({ name: user.name, email: user.email }); // Populate the form with the selected user's details
  };

  // Handle deleting a user
  const handleDelete = async (id: number) => {
    await axios.delete(`/api/users/${id}`); // DELETE request to remove the user
    fetchUsers(); // Refresh user list after deletion
  };

  // Render the component
  return (
    <Box p={2}>
      {/* Table to display the list of users */}
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>Name</TableCell>
              <TableCell>Email</TableCell>
              <TableCell>Actions</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {users.map((user) => (
              <TableRow key={user.id}>
                <TableCell>{user.id}</TableCell> {/* Display user ID */}
                <TableCell>{user.name}</TableCell> {/* Display user name */}
                <TableCell>{user.email}</TableCell> {/* Display user email */}
                <TableCell>
                  {/* Button to edit the user */}
                  <Button onClick={() => handleEdit(user)}>Edit</Button>
                  {/* Button to delete the user */}
                  <Button onClick={() => handleDelete(user.id)}>Delete</Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>

      {/* Form to create or update a user */}
      <Box mt={2}>
        <TextField
          label="Name"
          value={form.name} // Bind name input to form state
          onChange={(e) => setForm({ ...form, name: e.target.value })} // Update form state on change
          fullWidth
        />
        <TextField
          label="Email"
          value={form.email} // Bind email input to form state
          onChange={(e) => setForm({ ...form, email: e.target.value })} // Update form state on change
          fullWidth
          margin="normal"
        />
        {/* Button to save or create a user */}
        <Button onClick={handleSave}>
          {selectedUser ? "Save" : "Create"} {/* Change label based on context */}
        </Button>
      </Box>
    </Box>
  );
};

export default Home;
