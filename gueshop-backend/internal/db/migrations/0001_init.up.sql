-- AUTH MODULE

CREATE TABLE auth_profiles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE auth_permissions (
    code VARCHAR(10) PRIMARY KEY,
    description VARCHAR(255)
);

CREATE TABLE auth_profile_permissions (
    profile_id INTEGER NOT NULL REFERENCES auth_profiles(id) ON DELETE CASCADE,
    permission_code VARCHAR(10) NOT NULL REFERENCES auth_permissions(code) ON DELETE CASCADE,
    PRIMARY KEY (profile_id, permission_code)
);

CREATE TABLE auth_users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(500) NOT NULL,
    phone VARCHAR(50),
    profile_id INTEGER REFERENCES auth_profiles(id),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

-- CATALOG MODULE

CREATE TABLE cat_categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    parent_id UUID REFERENCES cat_categories(id) ON DELETE SET NULL,
    description TEXT
);

CREATE TABLE cat_suppliers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    contact_name VARCHAR(255),
    email VARCHAR(255)
);

CREATE TABLE cat_products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(500) NOT NULL,
    description TEXT,
    price NUMERIC(10,2) NOT NULL,
    stock_quantity INTEGER NOT NULL DEFAULT 0,
    category_id UUID REFERENCES cat_categories(id),
    supplier_id UUID REFERENCES cat_suppliers(id),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

-- CART MODULE

CREATE TABLE cart_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES auth_users(id),
    product_id UUID REFERENCES cat_products(id),
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);


-- ORDER MODULE

CREATE TABLE ord_orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES auth_users(id),
    total_amount NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

CREATE TABLE ord_order_items (
    id UUID PRIMARY KEY,
    order_id UUID REFERENCES ord_orders(id) ON DELETE CASCADE,
    product_id UUID REFERENCES cat_products(id),
    quantity INTEGER NOT NULL,
    unit_price NUMERIC(10,2) NOT NULL
);

CREATE TABLE ord_status (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255) -- Ej: 'pending', 'paid', 'shipped', 'cancelled'
);

CREATE TABLE ord_status_history
 (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID REFERENCES ord_orders(id),
    status_id INTEGER REFERENCES ord_status(id),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);


-- PAYMENT MODULE

CREATE TABLE payment_transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID REFERENCES ord_orders(id),
    payment_method VARCHAR(255) NOT NULL, -- Ej: 'webpay', 'mercadopago'
    status VARCHAR(255) NOT NULL,         -- Ej: 'success', 'failed', 'pending'
    amount NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

-- SHIPPING MODULE
CREATE TABLE ship_addresses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES auth_users(id),
    full_name VARCHAR(500) NOT NULL,
    address_line VARCHAR(500) NOT NULL,
    city VARCHAR(255) NOT NULL,
    region VARCHAR(500),
    postal_code VARCHAR(255),
    country VARCHAR(255) NOT NULL,
    phone VARCHAR(50)
);

CREATE TABLE ship_status_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID REFERENCES ord_orders(id),
    status VARCHAR(255) NOT NULL, -- Ej: 'pending', 'dispatched', 'in_transit', 'delivered'
    tracking_number VARCHAR(500),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

-- ADMIN MODULE

CREATE TABLE admin_activity_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES auth_users(id),
    action VARCHAR(255) NOT NULL,
    metadata JSONB,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

-- CUSTOMER MODULE

CREATE TABLE customer_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES auth_users(id),
    birthdate DATE,
    gender VARCHAR(100),
    loyalty_points INTEGER DEFAULT 0
);

-- NOTIFICATION MODULE

CREATE TABLE notification_emails (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES auth_users(id),
    type VARCHAR(255) NOT NULL, -- Ej: 'order_confirmed', 'password_reset'
    subject VARCHAR(255),
    body TEXT,
    sent_at TIMESTAMP WITHOUT TIME ZONE
);

-- CMS MODULE

CREATE TABLE cms_pages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    content TEXT,
    published BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);