export const DOTS_PATTERN = `
    <svg class="w-full h-full" xmlns="http://www.w3.org/2000/svg">
        <defs>
            <pattern id="dots" width="28" height="28" patternUnits="userSpaceOnUse">
            <circle cx="14" cy="14" r="1" fill="var(--color-brand-200)" opacity="0.35"/>
            </pattern>
        </defs>
        <rect width="100%" height="100%" fill="url(#dots)"/>
    </svg>
`;

export const GRID_PATTERN = `
    <svg class="w-full h-full" xmlns="http://www.w3.org/2000/svg">
        <defs>
        <pattern id="grid" width="40" height="40" patternUnits="userSpaceOnUse">
            <path d="M 40 0 L 0 0 0 40" fill="none" stroke="rgba(255,255,255,0.05)" stroke-width="0.8"/>
        </pattern>
        </defs>
        <rect width="100%" height="100%" fill="url(#grid)" />
    </svg>
`;