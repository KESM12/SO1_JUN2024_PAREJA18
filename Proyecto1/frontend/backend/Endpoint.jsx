export function ram() {
    return fetch(`/ram`, {
        headers: {'Content-Type': 'application/json'},
        method: 'GET',
    });
}