interface MenuItem {
  title: string
  icon: string
  activated: boolean
}

export interface Author {
  user_id: number
  full_name: string
  avatar_url: string
}

export interface Article {
  id: number
  createdTime: string
  updatedTime: string
  title: string
  author: Author
  content: string
  comments: number
  likes: number
  reads: number
}

export interface Comment {
  id: number
  createdTime: string
  content: string
  author: Author
  articleId: number
  rootCommentId: number
  replies: number
}

export interface RegisterForm {
  username:string
  password:string
  confirm_password:string
  full_name:string
}

export interface LoginForm {
  username:string
  password:string
}

export interface Profile {
  full_name: string
  bio: string
}

export interface Category {
  id: number
  title: string
  icon: string
  activated: boolean
}

export interface SubCategory {
  id:number
  title:string
  parentId:number
  activated:boolean
}

export interface Config {
  pagesize: number
  pagenum: number
}

export interface ArticlePostData {
  title: string
  content: string
  category_id: number
  content_type: string
}

export interface ArticleReciveData {
  id: number
  created_at: string
  updated_at: string
  title: string
  content: string
  content_type: string
  img: string
  comment_count: number
  read_count: number
  likes: number
  category: {
    category_id: number
    category_name: string
  }
  author: {
    user_id: number
    full_name: string
    avatar_url: string
  }
}

export interface SelfProfile {
  id: number
  username: string
  full_name: string
  register_date: string
  bio: string
  role: number
  avatar_url: string
}

export interface CommentPostData {
  content: string
  article_id: number
}

export interface CommentReciveData {
  content: string
  create_at: string
  id: number
  likes: number
  replied_user: number
  replies: ReplyReciveData[]
  total_replies: number
  user: {
    user_id: number
    full_name: string
    avatar_url: string
  }
}

export interface ReplyReciveData {
  id: number
  create_at: string
  content: string
  user: {
    user_id: number
    full_name: string
    avatar_url: string
  }
  likes: number
  replied_user: {
    user_id: number
    full_name: string
    avatar_url: string
  }
  replies:object[]
  total_replies: number
}

export interface ReplyPostData {
  content: string
  article_id: number
  comment_id: number
}