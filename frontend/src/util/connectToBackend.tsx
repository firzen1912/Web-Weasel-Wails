import { RunBackend } from '../../wailsjs/go/backend/Backend';

const connectToBackend = () => {
    // Connect to the backend here
    RunBackend().then((backend) => {
        console.log('Connected to backend');
    });
} 

export default connectToBackend;