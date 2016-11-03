CREATE ROLE bloom_rw WITH PASSWORD 'bloom_rwpassword';
ALTER ROLE bloom_rw WITH LOGIN;
GRANT CONNECT ON DATABASE bloomdb TO bloom_rw;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO bloom_rw;
GRANT ALL ON users TO bloom_rw;
GRANT ALL ON logins TO bloom_rw;
GRANT ALL ON role_perm TO bloom_rw;
GRANT ALL ON role_t TO bloom_rw;
GRANT ALL ON roles TO bloom_rw;
GRANT ALL ON perm TO bloom_rw;
GRANT ALL ON project TO bloom_rw;
GRANT ALL ON crosses TO bloom_rw;
GRANT ALL ON cross_parent TO bloom_rw;
GRANT ALL ON specimen TO bloom_rw;
GRANT ALL ON specimen_trait TO bloom_rw;
GRANT ALL ON trait TO bloom_rw;
GRANT ALL ON trait_t TO bloom_rw;
GRANT ALL ON mail TO bloom_rw;
