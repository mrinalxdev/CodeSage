import socket
import threading

connected_clients = []

def handle_client(client_socket, client_Address) :
    global connected_clients

    while True :
        try :
            data = client_socket.recv(4096).decode()
            if not data: 
                break
            print(f"Recieved data from {client_Address[0]}:{client_Address[1]}", data)

            if data.startswith("!reverse"):
                reversed_data = data[8:][::-1]
                client_socket.send(reversed_data.encode())
                print("Reversed data echoed back to client.")
            else :
                client_socket.send(data.encode())
                print("Data echoed back to client.")

        except Exception as e:
            print("An error occured while handling client :", e)
            break
    print(f"Client disconnected : {client_Address[0]}:{client_Address[1]}")
    client_socket.close()
    connected_clients.remove(client_socket)

def broadcast_message(message):
    global connected_clients

    for client_socket in connected_clients:
        try : 
            client_socket.send(message.encode())
        except Exception as e :
            print("An error occured while broadcasting message to client", e)

def main():
    SERVER_HOST= "127.0.0.1"
    SERVER_PORT = 8080

    try :
        server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        server_socket.bind((SERVER_HOST, SERVER_PORT))

        print(f"Server listening on {SERVER_HOST} : {SERVER_PORT}")

        while True:
            client_socket, client_Address = server_socket.accept()
            connected_clients.append(client_socket)
            print(f"Client connected from {client_Address[0]}:{client_Address[1]}")

            print(f"Number of connected clients : {len(connected_clients)}")

            client_handler = threading.Thread(target=handle_client, args = (client_socket, client_Address))
            client_handler.start()

            choice = input("Do you want to continue accepting connections ?? (Y/N): ")
            if choice.lower() != 'Y' :
                print('Server Shutting down....')
                break

    except socket.error as e:
        print("Socket err occurred :", e)
    
    except KeyboardInterrupt: 
        print("Server interrupted. Shutting down ....")
    
    except Exception as e:
        print("An error occurred : ", e)
    
    finally :
        # Closing all connected clients sockets
        for client_socket in connected_clients :
            client_socket.close()
        
        server_socket.close()


if __name__ == '__main__' :
    main()