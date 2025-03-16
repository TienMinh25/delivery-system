import {
  Box,
  Heading,
  Text,
  Button,
  Image,
  VStack,
  Container,
} from '@chakra-ui/react';
import { Link as RouterLink } from 'react-router-dom';

const NotFound = () => {
  return (
    <Container maxW='container.xl' py={20}>
      <VStack spacing={8} textAlign='center'>
        <Image
          src='/src/assets/images/not-found.svg'
          alt='Không tìm thấy trang'
          maxW='300px'
          fallbackSrc='https://via.placeholder.com/300x300?text=404'
        />

        <Heading as='h1' size='2xl' color='brand.600'>
          404
        </Heading>

        <Heading as='h2' size='xl' mb={4}>
          Không tìm thấy trang
        </Heading>

        <Text fontSize='lg' color='gray.600' maxW='md'>
          Trang bạn đang tìm kiếm có thể đã bị xóa, đổi tên hoặc tạm thời không
          khả dụng.
        </Text>

        <Box pt={4}>
          <Button
            as={RouterLink}
            to='/'
            colorScheme='brand'
            size='lg'
            rightIcon={
              <Box as='span' ml={2}>
                →
              </Box>
            }
          >
            Trở về trang chủ
          </Button>
        </Box>
      </VStack>
    </Container>
  );
};

export default NotFound;
