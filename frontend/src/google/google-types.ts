declare global {
    interface Window {
        google: {
            accounts: {
                id: {
                    initialize: (options: unknown) => void;
                    renderButton: (parent: HTMLElement, options: unknown) => void;
                    prompt: () => void;
                };
            };
        };
    }
}

export interface CredentialResponse {
    credential: string;
    select_by: 'auto' | 'user' | 'user_1tap' | 'user_2tap';
    clientId: string;
}

