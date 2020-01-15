CREATE FUNCTION jsonb_merge(orig JSONB, delta JSONB)
RETURNS JSONB LANGUAGE SQL AS $$
    SELECT
        jsonb_object_agg(
            COALESCE(keyOrig, keyDelta),
            CASE
                WHEN valOrig ISNULL THEN valDelta
                WHEN valDelta ISNULL THEN valOrig
                WHEN (jsonb_typeof(valOrig) != 'object' OR jsonb_typeof(valDelta) != 'object') THEN valDelta
                ELSE jsonb_merge(valOrig, valDelta)
            END
        )
    FROM jsonb_each(orig) AS rowOrig(keyOrig, valOrig)
    FULL JOIN jsonb_each(delta) AS rowDelta(keyDelta, valDelta) ON
        keyOrig = keyDelta
$$;
