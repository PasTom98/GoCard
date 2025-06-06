/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./index.html",
        "./src/**/*.{vue,js,ts,jsx,tsx}",
    ],
    darkMode: 'class',
    theme: {
        extend: {
            colors: {
                'vt-c-white': '#ffffff',
                'vt-c-white-soft': '#f8f8f8',
                'vt-c-white-mute': '#f2f2f2',
                'vt-c-black': '#181818',
                'vt-c-black-soft': '#222222',
                'vt-c-black-mute': '#282828',
                'vt-c-indigo': '#2c3e50',
                'vt-c-divider-light-1': 'rgba(60, 60, 60, 0.29)',
                'vt-c-divider-light-2': 'rgba(60, 60, 60, 0.12)',
                'vt-c-divider-dark-1': 'rgba(84, 84, 84, 0.65)',
                'vt-c-divider-dark-2': 'rgba(84, 84, 84, 0.48)',
                'vt-c-text-light-1': '#2c3e50',
                'vt-c-text-light-2': 'rgba(60, 60, 60, 0.66)',
                'vt-c-text-dark-1': '#ffffff',
                'vt-c-text-dark-2': 'rgba(235, 235, 235, 0.64)',
            },
            fontFamily: {
                sans: [
                    'Inter',
                    '-apple-system',
                    'BlinkMacSystemFont',
                    'Segoe UI',
                    'Roboto',
                    'Oxygen',
                    'Ubuntu',
                    'Cantarell',
                    'Fira Sans',
                    'Droid Sans',
                    'Helvetica Neue',
                    'sans-serif',
                ],
            },
            boxShadow: {
                'card': '0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)',
                'card-hover': '0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)',
            },
        },
    },
    plugins: [],
};
