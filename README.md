# AI-Enhanced Social Media Platform

A modern, AI-powered social media platform that combines artificial intelligence image generation with traditional social media features. Built with Go backend and React frontend, featuring Elasticsearch for search capabilities and Google Cloud Storage for media management.

## 🚀 Features

- **AI Image Generation**: Create stunning images using OpenAI's DALL-E 3 model
- **User Authentication**: Secure JWT-based authentication system
- **Media Management**: Upload, store, and organize images and videos
- **Advanced Search**: Elasticsearch-powered search across posts and users
- **Responsive Design**: Modern Material-UI based interface
- **Photo Gallery**: Interactive photo viewing with lightbox functionality
- **Cloud Storage**: Google Cloud Storage integration for scalable media storage

## 🏗️ Architecture

### Backend (Go)

- **Framework**: Go with Gorilla Mux router
- **Authentication**: JWT middleware with Auth0
- **Database**: Elasticsearch for search and indexing
- **Storage**: Google Cloud Storage for media files
- **Deployment**: Google App Engine (Flexible environment)

### Frontend (React)

- **Framework**: React 19 with React Router
- **UI Library**: Material-UI (MUI) components
- **Styling**: Styled-components and CSS-in-JS
- **State Management**: React hooks and localStorage
- **HTTP Client**: Axios for API communication

## 📁 Project Structure

```
ai-enhanced-social-media-platform/
├── backend/                 # Go backend application
│   ├── backend/            # Core backend services
│   │   ├── elasticsearch.go # Elasticsearch client
│   │   └── gcs.go         # Google Cloud Storage client
│   ├── constants/          # Configuration constants
│   ├── handler/            # HTTP request handlers
│   ├── model/              # Data models
│   ├── service/            # Business logic services
│   ├── app.yaml           # App Engine configuration
│   ├── go.mod             # Go dependencies
│   └── main.go            # Application entry point
├── frontend/               # React frontend application
│   ├── src/
│   │   ├── components/     # React components
│   │   ├── styles/         # CSS stylesheets
│   │   └── constants.js    # Frontend constants
│   ├── package.json        # Node.js dependencies
│   └── public/             # Static assets
└── README.md               # This file
```

## 🛠️ Prerequisites

### Backend Requirements

- Go 1.22+
- Elasticsearch instance
- Google Cloud Platform account
- Google Cloud Storage bucket

### Frontend Requirements

- Node.js 16+
- npm or yarn
- OpenAI API key

## 🚀 Getting Started

### Backend Setup

1. **Navigate to backend directory**:

   ```bash
   cd backend
   ```

2. **Install Go dependencies**:

   ```bash
   go mod download
   ```

3. **Configure environment variables**:

   - Set Elasticsearch connection details in `constants/constants.go`
   - Configure Google Cloud credentials
   - Set GCS bucket name

4. **Run the application**:

   ```bash
   go run main.go
   ```

5. **Deploy to Google App Engine**:
   ```bash
   gcloud app deploy
   ```

### Frontend Setup

1. **Navigate to frontend directory**:

   ```bash
   cd frontend
   ```

2. **Install dependencies**:

   ```bash
   npm install
   ```

3. **Set environment variables**:
   Create a `.env` file with:

   ```
   REACT_APP_OPENAI_KEY=your_openai_api_key
   ```

4. **Start development server**:

   ```bash
   npm start
   ```

5. **Build for production**:
   ```bash
   npm run build
   ```

## 🔧 Configuration

### Backend Configuration

- **Elasticsearch**: Configure connection URL, username, and password
- **Google Cloud Storage**: Set bucket name and ensure proper authentication
- **JWT**: Configure signing key for authentication

### Frontend Configuration

- **API Base URL**: Update `BASE_URL` in `constants.js` to match your backend deployment
- **OpenAI API**: Set your OpenAI API key in environment variables

## 📡 API Endpoints

### Authentication

- `POST /signup` - User registration
- `POST /signin` - User login

### Content Management

- `POST /upload` - Upload media files (requires authentication)
- `GET /search` - Search posts and users (requires authentication)

## 🔐 Security Features

- JWT-based authentication
- CORS configuration
- Secure file upload handling
- Protected API endpoints

## 🌟 Key Features Explained

### AI Image Generation

The platform integrates with OpenAI's DALL-E 3 model to generate high-quality images based on text descriptions. Users can:

- Enter detailed prompts for image generation
- Preview generated images in a responsive gallery
- Upload AI-generated images to their collection

### Media Management

- **Upload**: Support for images and videos
- **Storage**: Google Cloud Storage integration for scalable storage
- **Organization**: Elasticsearch indexing for fast search and retrieval

### User Experience

- **Responsive Design**: Mobile-first approach with Material-UI
- **Photo Gallery**: Interactive viewing with lightbox functionality
- **Search**: Advanced search capabilities across content and users

## 🚀 Deployment

### Backend Deployment

The backend is configured for Google App Engine deployment with:

- Automatic scaling (1-2 instances)
- Ubuntu 22.04 runtime
- Go 1.21+ runtime

### Frontend Deployment

The frontend can be deployed to any static hosting service:

- Netlify
- Vercel
- Google Cloud Storage (static website)
- AWS S3 + CloudFront

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Support

For support and questions:

- Check the existing issues
- Create a new issue with detailed description
- Contact the development team

## 🔮 Future Enhancements

- Real-time notifications
- Social features (likes, comments, sharing)
- Advanced AI filters and effects
- Mobile application
- Video generation capabilities
- Community features and groups

---

**Built with ❤️ using Go, React, and AI LLMs**
