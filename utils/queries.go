package utils

const (
	GetAllProductActive = `
		SELECT 
			c.id AS id_category,
			c.name AS name_category,
			c.description AS deskripsi,
			COALESCE(json_agg(
				json_build_object(
					'id', p.id,
					'category_id', p.category_id,
					'name', p.name,
					'sku', p.sku,
					'description', p.description,
					'price', p.price,
					'promo_price', p.promo_price,
					'stock', p.stock,
					'weight', p.weight,
					'image_url', p.image_url,
					'is_active', p.is_active,
					'created_at', p.created_at,
					'updated_at', p.updated_at
				)
			) FILTER (WHERE p.id IS NOT NULL), '[]') AS product
		FROM 
			product_categories c
		LEFT JOIN 
			products p ON p.category_id = c.id
		GROUP BY 
			c.id, c.name, c.description;
	`
)
