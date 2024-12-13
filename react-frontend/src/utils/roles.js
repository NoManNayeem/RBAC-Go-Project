export const getRole = () => {
    // Simulate fetching the role; replace with real user role retrieval logic
    // Example roles: "user", "admin"
    return "user"; // Change this to "admin" to simulate an admin
};

export const hasPermission = (requiredPermission) => {
    const rolePermissions = {
        user: ["read"],
        admin: ["read", "write", "delete"],
    };

    const role = getRole();
    return rolePermissions[role]?.includes(requiredPermission);
};
