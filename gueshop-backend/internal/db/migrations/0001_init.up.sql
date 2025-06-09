-- AUTH MODULE

-- Table: auth_profiles
CREATE TABLE auth_profiles (
    id UUID PRIMARY KEY,
    code TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL
);

-- Table: auth_permissions
CREATE TABLE auth_permissions (
    id UUID PRIMARY KEY,
    code TEXT UNIQUE NOT NULL,
    description TEXT
);

-- Table: auth_profile_permissions
CREATE TABLE auth_profile_permissions (
    id UUID PRIMARY KEY,
    profile_id UUID NOT NULL REFERENCES auth_profiles(id) ON DELETE CASCADE,
    permission_id UUID NOT NULL REFERENCES auth_permissions(id) ON DELETE CASCADE,
    UNIQUE (profile_id, permission_id)
);

-- Table: auth_users
CREATE TABLE auth_users (
    id UUID PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    phone TEXT,
    profile_id UUID REFERENCES auth_profiles(id),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

-- CATALOG MODULE
-- Table: catalog_categories
CREATE TABLE catalog_categories (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    parent_id UUID REFERENCES catalog_categories(id) ON DELETE SET NULL,
    description TEXT
);

-- Table: catalog_suppliers
CREATE TABLE catalog_suppliers (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    contact_name TEXT,
    email TEXT
);

-- Table: catalog_products
CREATE TABLE catalog_products (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    price NUMERIC(10,2) NOT NULL,
    stock_quantity INTEGER NOT NULL DEFAULT 0,
    category_id UUID REFERENCES catalog_categories(id),
    supplier_id UUID REFERENCES catalog_suppliers(id),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);


-- CART MODULE
-- Table: cart_items
CREATE TABLE cart_items (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES auth_users(id),
    product_id UUID REFERENCES catalog_products(id),
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);


-- ORDER MODULE
-- Table: order_orders
CREATE TABLE order_orders (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES auth_users(id),
    status TEXT NOT NULL, -- Ej: 'pending', 'paid', 'shipped', 'cancelled'
    total_amount NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

-- Table: order_order_items
CREATE TABLE order_order_items (
    id UUID PRIMARY KEY,
    order_id UUID REFERENCES order_orders(id) ON DELETE CASCADE,
    product_id UUID REFERENCES catalog_products(id),
    quantity INTEGER NOT NULL,
    unit_price NUMERIC(10,2) NOT NULL
);


-- PAYMENT MODULE
-- Table: payment_transactions
CREATE TABLE payment_transactions (
    id UUID PRIMARY KEY,
    order_id UUID REFERENCES order_orders(id),
    payment_method TEXT NOT NULL, -- Ej: 'webpay', 'mercadopago'
    status TEXT NOT NULL,         -- Ej: 'success', 'failed', 'pending'
    amount NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);


-- SHIPPING MODULE
-- Table: shipping_addresses
CREATE TABLE shipping_addresses (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES auth_users(id),
    full_name TEXT NOT NULL,
    address_line1 TEXT NOT NULL,
    address_line2 TEXT,
    city TEXT NOT NULL,
    region TEXT,
    postal_code TEXT,
    country TEXT NOT NULL,
    phone TEXT
);

-- Table: shipping_status
CREATE TABLE shipping_status (
    id UUID PRIMARY KEY,
    order_id UUID REFERENCES order_orders(id),
    status TEXT NOT NULL, -- Ej: 'pending', 'dispatched', 'in_transit', 'delivered'
    tracking_number TEXT,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);


-- ADMIN MODULE
-- Table: admin_activity_logs
CREATE TABLE admin_activity_logs (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES auth_users(id),
    action TEXT NOT NULL,
    metadata JSONB,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);


-- CUSTOMER MODULE
-- Table: customer_profiles
CREATE TABLE customer_profiles (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES auth_users(id),
    birthdate DATE,
    gender TEXT,
    loyalty_points INTEGER DEFAULT 0
);


-- NOTIFICATION MODULE
-- Table: notification_emails
CREATE TABLE notification_emails (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES auth_users(id),
    type TEXT NOT NULL, -- Ej: 'order_confirmed', 'password_reset'
    subject TEXT,
    body TEXT,
    sent_at TIMESTAMP WITHOUT TIME ZONE
);


-- CMS MODULE
-- Table: cms_pages
CREATE TABLE cms_pages (
    id UUID PRIMARY KEY,
    slug TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    content TEXT,
    published BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);
