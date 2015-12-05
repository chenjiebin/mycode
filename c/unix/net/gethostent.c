
/*
	#include <netdb.h>
	struct hostent *gethostent(void);
		返回值：若成功，返回指针；若出错，返回NULL
		hostent
			struct hostent {
				char 	*h_name;
				char 	**h_aliases;
				int 	h_addrtype;
				int 	h_length;
				char	**h_addr_list;
				...
			}
	void sethostent(int stayopen);
	void endhostent(void);


	#include <netdb.h>
	struct netent *getnetbyaddr (uint32_t net, int type);
	struct netent *getnetbyname(const char *name);
	struct netent *getnetent(void);
		3个函数返回值：若成功返回指针；若出错返回NULL
		netent
			struct netent {
				char 		*n_name
				char 		**n_aliases;
				int 		n_addrtype;
				uint32_t	n_net;
				...
			}
	void setnetent(int stayopen);
	void endnetent(void);


	#include <netdb.h>
	struct protoent *getprotobyname(const char *name);
	struct protoent *getprotobynumber(int proto);
	struct protoent *getprotoent(void);
		3个函数返回值：若成功返回指针；若出错返回NULL
		protoent
			struct {
				char	*p_name;
				char	**p_aliases;
				int 	p_proto;
				...
			}
	void setprotoent(int stayopen);
	void endprotoent(void);


	#include <netdb.h>
	struct servent *getservbyname(const char *name, const char *proto);
	struct servent *getserbyport(int port, const char *proto);
	struct servent *getservent(void);
		3个函数返回值：若成功返回指针，若出错返回NULL
		struct servent {
			char	*s_name;
			char 	**s_aliases;
			int		s_port;
			char	*s_proto;
			...
		}
	void setservent(int stayopen);
	void endservent(void);


	#include <sys/socket.h>
	#include <netdb.h>
	int getaddrinfo(const char *restrict host,
					const char *restrict service,
					const struct addrinfo *restrict hint,
					struct addrinfo **restrict res);
		返回值：若成功返回0；若出错返回非0错误码
	voidfreeaddrinfo(struct addrinfo *ai);
		addrinfo
			struct addrinfo {
				int 		ai_flags;
				int			ai_family;
				int			ai_socktype;
				int			ai_protocol;
				socklen_t	ai_addrlen;
				struct sockaddr		*ai_addr;
				char				*ai_canonname;
				struct addrinfo		*ai_next;
				...
			}


*/





























