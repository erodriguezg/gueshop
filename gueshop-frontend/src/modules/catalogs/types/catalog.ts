export interface Category {
  name: string;
  description: string;
  imageUrl?: string;
  slug: string;
}

export interface Product {
  id: string;
  name: string;
  shortDescription: string;
  imageUrl?: string;
  price: number;
}
