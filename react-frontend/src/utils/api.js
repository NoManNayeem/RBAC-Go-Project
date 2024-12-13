const API_URL = process.env.REACT_APP_API_URL;

export const fetchPublicData = async () => {
    try {
        const response = await fetch(`${API_URL}/public`);
        if (!response.ok) {
            throw new Error('Failed to fetch public data');
        }
        return await response.text();
    } catch (error) {
        console.error(error);
        return null;
    }
};

export const fetchAdminData = async () => {
    try {
        const response = await fetch(`${API_URL}/admin`, {
            headers: {
                Role: "admin", // Replace with dynamic role if needed
            },
        });
        if (!response.ok) {
            throw new Error('Failed to fetch admin data');
        }
        return await response.text();
    } catch (error) {
        console.error(error);
        return null;
    }
};
