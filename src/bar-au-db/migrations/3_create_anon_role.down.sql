REVOKE usage ON SCHEMA public
FROM
    anon;

REVOKE SELECT ON public.bars
FROM
    anon;

DROP ROLE anon;

